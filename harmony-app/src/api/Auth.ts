import { ReqHelper, Auth } from './ReqHelper';

interface IRegisterResponse {
	session: string;
}

interface ILoginResponse {
	session: string;
}

export class AuthAPI {
	static endPoint = `http://localhost:2289/api/v1`;
	static token = localStorage.getItem('token');

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
			Auth.TOKEN
		);
	}

	static async login(email: string, password: string) {
		return ReqHelper.post<ILoginResponse>(`${this.endPoint}/login`, null, { email, password }, Auth.TOKEN);
	}
}

//@ts-ignore
window.AuthConnection = AuthAPI;
