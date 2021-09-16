// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package auth_test

import (
	"strings"
	"testing"

	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/harmony-development/legato/logger"
	"github.com/harmony-development/legato/server"
)

func contains(s string, ss []string) bool {
	for _, it := range ss {
		if it == s {
			return true
		}
	}

	return false
}

func test(t *testing.T, s string, i int, fn func(t *testing.T, i int)) {
	t.Helper()
	t.Logf("%sTesting %s", strings.Repeat("\t", i), s)
	fn(t, i+1)
}

func beginAuth(client authv1.HTTPTestAuthServiceClient, authid *string) func(t *testing.T, i int) {
	return func(t *testing.T, i int) {
		t.Helper()

		resp, err := client.BeginAuth(&authv1.BeginAuthRequest{})
		if err != nil {
			t.Fatalf("error: %s", err)
		}

		*authid = resp.AuthId
	}
}

func firstAuthStep(client authv1.HTTPTestAuthServiceClient, authid, is string) func(t *testing.T, i int) {
	return func(t *testing.T, i int) {
		t.Helper()

		resp, err := client.NextStep(&authv1.NextStepRequest{
			AuthId: authid,
		})
		if err != nil {
			t.Fatalf("error: %s", err)
		}

		choice, ok := resp.Step.Step.(*authv1.AuthStep_Choice_)
		if !ok {
			t.Fatalf("first thing wasn't choice")
		}

		if !contains(is, choice.Choice.Options) {
			t.Fatalf("no '%s' in options", is)
		}
	}
}

func formAuthStep(client authv1.HTTPTestAuthServiceClient, authid, step string) func(t *testing.T, i int) {
	return func(t *testing.T, i int) {
		t.Helper()

		resp, err := client.NextStep(&authv1.NextStepRequest{
			AuthId: authid,
			Step: &authv1.NextStepRequest_Choice_{
				Choice: &authv1.NextStepRequest_Choice{
					Choice: step,
				},
			},
		})
		if err != nil {
			t.Fatalf("error: %s", err)
		}

		_, ok := resp.Step.Step.(*authv1.AuthStep_Form_)
		if !ok {
			t.Fatalf("step wasn't form")
		}
	}
}

func register(
	client authv1.HTTPTestAuthServiceClient,
	authid,
	username,
	email,
	password string,
) func(t *testing.T, i int) {
	return func(t *testing.T, i int) {
		t.Helper()

		resp, err := client.NextStep(&authv1.NextStepRequest{
			AuthId: authid,
			Step: &authv1.NextStepRequest_Form_{
				Form: &authv1.NextStepRequest_Form{
					Fields: []*authv1.NextStepRequest_FormFields{
						{
							Field: &authv1.NextStepRequest_FormFields_String_{
								String_: email,
							},
						},
						{
							Field: &authv1.NextStepRequest_FormFields_String_{
								String_: username,
							},
						},
						{
							Field: &authv1.NextStepRequest_FormFields_Bytes{
								Bytes: []byte(password),
							},
						},
					},
				},
			},
		})
		if err != nil {
			t.Fatalf("error: %s", err)
		}

		session, ok := resp.Step.Step.(*authv1.AuthStep_Session)
		if !ok {
			t.Fatalf("register wasn't session, got %+v", resp.Step.Step)
		}

		_ = session
	}
}

func login(client authv1.HTTPTestAuthServiceClient, authid, email, password string) func(t *testing.T, i int) {
	return func(t *testing.T, i int) {
		t.Helper()

		resp, err := client.NextStep(&authv1.NextStepRequest{
			AuthId: authid,
			Step: &authv1.NextStepRequest_Form_{
				Form: &authv1.NextStepRequest_Form{
					Fields: []*authv1.NextStepRequest_FormFields{
						{
							Field: &authv1.NextStepRequest_FormFields_String_{
								String_: email,
							},
						},
						{
							Field: &authv1.NextStepRequest_FormFields_Bytes{
								Bytes: []byte(password),
							},
						},
					},
				},
			},
		})
		if err != nil {
			t.Fatalf("error: %s", err)
		}

		session, ok := resp.Step.Step.(*authv1.AuthStep_Session)
		if !ok {
			t.Fatalf("login wasn't session, got %+v", resp.Step.Step)
		}

		_ = session
	}
}

// nolint
// Integration tests cannot be parallelized
func TestAuth(t *testing.T) {
	l := logger.NewNoop()

	serv, err := server.New(l)
	if err != nil {
		t.Fatal(err)
	}

	client := authv1.HTTPTestAuthServiceClient{}
	client.Client = serv

	test(t, "client auth", 0, func(t *testing.T, i int) {
		var authid string
		const (
			username = "kili-test"
			email    = "uhh@eee@aaa"
			password = "kala-test"
		)

		test(t, "begin auth", i, beginAuth(client, &authid))
		test(t, "first auth step", i, firstAuthStep(client, authid, "register"))
		test(t, "get register form", i, formAuthStep(client, authid, "register"))

		test(t, "register account", i, register(client, authid, username, email, password))

		test(t, "begin auth again", i, beginAuth(client, &authid))
		test(t, "first auth step again", i, firstAuthStep(client, authid, "login"))
		test(t, "get login form", i, formAuthStep(client, authid, "login"))

		test(t, "login account", i, login(client, authid, email, password))
	})
}
