import React from 'react';
import { Switch, Route, useHistory } from 'react-router';
import { BrowserRouter } from 'react-router-dom';
import { CssBaseline, createMuiTheme, Button } from '@material-ui/core';
import { ThemeProvider } from '@material-ui/core/styles';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer, toast, cssTransition } from 'react-toastify';
import { useSelector } from 'react-redux';

import { IState } from '../types/redux';
import HarmonySocket from '../socket/socket';

import { App } from './App/App';
import { Entry } from './Entry/Entry';
import { useRootStyles } from './RootStyle';
import './Root.css';
import { InviteHandler } from './InviteHandler/HandleInvite';
import { HarmonyDark } from './App/HarmonyColor';
import { useSocketHandler } from './SocketHandler';

export const harmonySocket = new HarmonySocket();

const Zoom = cssTransition({
	enter: 'zoomIn',
	exit: 'slideOut',
	duration: 200,
});

const RootWithRouter = (): JSX.Element => {
	useSocketHandler(harmonySocket, useHistory());

	return (
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
	);
};

const Root = (): JSX.Element => {
	useRootStyles();

	return (
		<>
			<CssBaseline />
			<ToastContainer position="bottom-left" pauseOnFocusLoss={false} transition={Zoom} />
			<BrowserRouter>
				<RootWithRouter />
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
			background: {
				default: themeState.type === 'dark' ? HarmonyDark[600] : '#FFF',
				paper: themeState.type === 'dark' ? HarmonyDark[500] : '#FFF',
			},
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
