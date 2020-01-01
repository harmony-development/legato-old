import React, { useEffect } from 'react';
import { Switch, Route } from 'react-router';
import { BrowserRouter } from 'react-router-dom';
import { CssBaseline, createMuiTheme, Button } from '@material-ui/core';
import { ThemeProvider } from '@material-ui/core/styles';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer, toast } from 'react-toastify';
import { useSelector } from 'react-redux';

import { IState } from '../types/redux';
import HarmonySocket from '../socket/socket';

import { App } from './App/App';
import { Entry } from './Entry/Entry';
import { useRootStyles } from './RootStyle';
import './Root.css';
import { InviteHandler } from './InviteHandler/HandleInvite';

import { store } from '../redux/store';
import { SetConnected } from '../redux/AppReducer';

export const harmonySocket = new HarmonySocket();
export let previouslyDisconnected = false;

const Root = (): JSX.Element => {
	useRootStyles();

	useEffect(() => {
		harmonySocket.events.addListener('close', () => {
			// lol plz no spahm
			if (!previouslyDisconnected) {
				toast.error('You have lost connection to the server');
				store.dispatch(SetConnected(true));
				previouslyDisconnected = true;
			}
			setTimeout(harmonySocket.connect, 3000);
		});
		harmonySocket.events.addListener('open', () => {
			if (previouslyDisconnected) toast.success('You have reconnected to the server');
			store.dispatch(SetConnected(true));
		});
		return (): void => {
			harmonySocket.events.removeAllListeners('close'); // cleanup all socket events registered here
			harmonySocket.events.removeAllListeners('open');
		};
	}, []);

	return (
		<>
			<CssBaseline />
			<ToastContainer />
			<BrowserRouter>
				<Switch>
					<Route exact path="/">
						<Entry />
					</Route>
					<Route exact path="/app/:selectedguildparam?/:selectedchannelparam?">
						<App />
					</Route>
					<Route exact path="/invite/:invitecode">
						<InviteHandler />
					</Route>
					<Route exact path="/bruh">
						<Button
							onClick={(): void => {
								toast.info('GET BRUHED ON KID');
							}}
						>
							Bruh Button
						</Button>
					</Route>
				</Switch>
			</BrowserRouter>
		</>
	);
};

const ThemedRoot = (): JSX.Element => {
	const themeState = useSelector((state: IState) => state.theme);
	const theme = createMuiTheme({
		palette: {
			primary: themeState.primary,
			secondary: themeState.secondary,
			type: themeState.type,
		},
	});

	return (
		<div className="root">
			<ThemeProvider theme={theme}>
				<Root />
			</ThemeProvider>
		</div>
	);
};

export default ThemedRoot;
