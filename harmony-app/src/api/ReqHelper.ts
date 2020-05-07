import i18n from '../i18n/i18n';

export enum Auth {
	NONE,
	TOKEN,
}

type GetParams = {
	[key: string]: any;
} | null;

export class ReqHelper {
	static authToken = localStorage.getItem('token') || undefined;
	static tokenKey = 'token';

	static async getToken() {
		return localStorage.getItem('token');
	}

	static async refreshToken() {
		this.authToken = (await this.getToken()) || undefined;
	}

	static async get<T>(url: string, params: GetParams, auth: Auth = Auth.TOKEN) {
		return this.http<T>('GET', url, params, null, auth);
	}

	static async post<T>(url: string, args: GetParams, body: any, auth: Auth = Auth.TOKEN) {
		return await this.http<T>('POST', url, args, body, auth);
	}

	static async put<T>(url: string, args: GetParams, body: any, auth: Auth = Auth.TOKEN) {
		return await this.http<T>('PUT', url, args, body, auth);
	}

	static async patch<T>(url: string, args: GetParams, body: any, auth: Auth = Auth.TOKEN) {
		return await this.http<T>('PATCH', url, args, body, auth);
	}

	static async delete<T>(url: string, args: GetParams, auth: Auth = Auth.TOKEN) {
		return await this.http<T>('DELETE', url, args, null, auth);
	}

	static async http<T>(method: string, rawURL: string, params: GetParams, body: any | null, auth: Auth) {
		const headers = new Headers();
		const url = new URL(rawURL);
		if (params) {
			Object.keys(params).forEach(param => {
				url.searchParams.set(param, params[param]);
			});
		}
		if (auth === Auth.TOKEN && this.authToken) {
			headers.set('Authorization', this.authToken);
		}
		if (body) {
			body = JSON.stringify(body);
			headers.set('Content-Type', 'application/json');
		}
		let response: Response;
		try {
			response = await fetch(url.toString(), {
				body,
				headers,
				method,
			});
		} catch (e) {
			throw new Error(i18n.t('errors:network'));
		}
		let data: any;
		try {
			data = await response.json();
		} catch (e) {
			throw new Error(i18n.t(`errors:parse-error`));
		}
		if (data.message && !response.ok) {
			throw new Error(i18n.t(`errors:${data.message}`, data.fields));
		}
		return data as T;
	}
}
