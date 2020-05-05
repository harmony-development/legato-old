import React, { useEffect, lazy, Suspense } from 'react';
import { Switch, Route } from 'react-router';
import { useSelector } from 'react-redux';
import { CssBaseline, createMuiTheme } from '@material-ui/core';
import { ThemeProvider, Theme, makeStyles } from '@material-ui/core/styles';

import 'react-toastify/dist/ReactToastify.css';
import HarmonySocket from '../socket/socket';
import { IState } from '../types/redux';

import './Root.css';
import { Loading } from '../component/Loading';

import { HarmonyDark } from './App/HarmonyColor';

export const harmonySocket = new HarmonySocket();

const InvitePage = lazy(async () => ({
	default: (await import('./InviteHandler/HandleInvite')).InviteHandler,
}));

const EntryPage = lazy(async () => ({
	default: (await import('./Entry/Entry')).Entry,
}));

const AppPage = lazy(async () => ({
	default: (await import('./App/App')).App,
}));

const rootStyles = makeStyles((theme: Theme) => ({
	'@global': {
		'::-webkit-scrollbar': {
			width: '10px',
		},
		'::-webkit-scrollbar-thumb:hover': {
			backgroundColor: theme.palette.type === 'light' ? 'rgb(150, 150, 150)' : 'rgb(100, 100, 100)',
		},
		'::-webkit-scrollbar-track': {
			backgroundColor: theme.palette.type === 'light' ? 'rgb(245, 245, 245)' : 'rgb(46, 46, 46)',
		},
		'::-webkit-scrollbar-thumb': {
			backgroundColor: theme.palette.type === 'light' ? 'rgb(200, 200, 200)' : 'rgb(64, 64, 64)',
		},
		'::-webkit-scrollbar-corner': {
			backgroundColor: theme.palette.type === 'light' ? 'rgb(200, 200, 200)' : 'rgb(64, 64, 64)',
		},
		'*': {
			scrollbarColor: `${theme.palette.type === 'light' ? 'rgb(200, 200, 200)' : 'rgb(64, 64, 64)'} ${
				theme.palette.type === 'light' ? 'rgb(245, 245, 245)' : 'rgb(46, 46, 46)'
			}`,
		},
	},
}));

export const Root = React.memo(() => {
	const themeState = useSelector((state: IState) => state.theme);
	const theme = createMuiTheme({
		palette: {
			primary: themeState.primary,
			secondary: themeState.secondary,
			type: themeState.type,
			background: {
				default: themeState.type === 'dark' ? HarmonyDark[600] : '#FFF',
				paper: themeState.type === 'dark' ? HarmonyDark[500] : '#f7f7f7',
			},
		},
	});
	rootStyles();

	useEffect(() => {
		if (!localStorage.getItem('developerCodeExecution')) {
			localStorage.setItem(
				'developerCodeExecution',
				'WARNING : SETTING THIS VALUE TO TRUE COULD LEAD TO VULNERABILITIES'
			);
		}
	}, []);

	return (
		<ThemeProvider theme={theme}>
			<CssBaseline />
			<Suspense fallback={Loading()}>
				<Switch>
					<Route exact path="/">
						<EntryPage />
					</Route>
					<Route exact path="/app/:selectedguildparam?/:selectedchannelparam?">
						<AppPage />
					</Route>
					<Route exact path="/invite/:invitecode">
						<InvitePage />
					</Route>
				</Switch>
			</Suspense>
		</ThemeProvider>
	);
});
