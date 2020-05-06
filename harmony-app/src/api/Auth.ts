import { ReqHelper, Auth } from './ReqHelper';

export class AuthAPI {
	static endPoint = `http://localhost:2288/api/v1`;
	static token = localStorage.getItem('token');

	static async register() {
		try {
			const resp = await ReqHelper.get(`${this.endPoint}/register`, null, Auth.TOKEN);
			console.log(resp);
		} catch (err) {
			console.error(err);
		}
	}

	static async login() {
		fetch(`${this.endPoint}/login`, {
			method: 'POST',
		});
	}
}

//@ts-ignore
window.AuthConnection = AuthAPI;
