class AuthConnection {
	endPoint: string;
	token: string | undefined;

	constructor() {
		this.endPoint = `${process.env.REACT_APP_HARMONY_AUTH_HOST}/api/${process.env.REACT_APP_HARMONY_AUTH_VERSION}`;
		this.token = localStorage.getItem('token') || undefined;
	}

	register() {}

	login() {
		fetch(`${this.endPoint}/login`, {
			method: 'POST'
		})
	}
}

export const AuthServer = new AuthConnection();
