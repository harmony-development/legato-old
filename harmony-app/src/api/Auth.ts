import { ReqHelper, Auth } from './ReqHelper';

interface IRegisterResponse {
	session: string;
}

interface ILoginResponse {
	session: string;
}

type IListServersResponse = {
	name: string;
	host: string;
}[];

export class AuthAPI {
	static endPoint = `https://localhost:2289/api/v1`;

	static async getSession() {
		return localStorage.getItem('authsession');
	}

	static async register(email: string, username: string, password: string) {
		const passwordStats = {
			capital: 0,
			lowercase: 0,
			numbers: 0,
			special: 0,
		};

		[...password].forEach(char => {
			if (char === char.toUpperCase()) passwordStats.capital++;
			if (char === char.toLowerCase()) passwordStats.lowercase++;
			if (!isNaN(Number(char))) passwordStats.numbers++;
			else passwordStats.special++;
		});

		return ReqHelper.post<IRegisterResponse>(
			`${this.endPoint}/register`,
			null,
			{
				email,
				username,
				password,
				passwordStats,
			},
			Auth.AUTH_SESSION
		);
	}

	static async login(email: string, password: string) {
		return ReqHelper.post<ILoginResponse>(`${this.endPoint}/login`, null, { email, password }, Auth.AUTH_SESSION);
	}

	static async listServers() {
		return ReqHelper.post<IListServersResponse>(`${this.endPoint}/listservers`, null, null, Auth.AUTH_SESSION);
	}
}

//@ts-ignore
window.AuthConnection = AuthAPI;
