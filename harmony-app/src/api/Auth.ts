import { ReqHelper, Auth } from './ReqHelper';

interface IRegisterResponse {
	userid: string;
	token: string;
}

export class AuthAPI {
	static endPoint = `http://localhost:2288/api/v1`;
	static token = localStorage.getItem('token');

	static async register(email: string, username: string, password: string) {
		return ReqHelper.post<IRegisterResponse>(
			`${this.endPoint}/register`,
			null,
			{ email, username, password },
			Auth.TOKEN
		);
	}

	static async login() {
		fetch(`${this.endPoint}/login`, {
			method: 'POST',
		});
	}
}

//@ts-ignore
window.AuthConnection = AuthAPI;
