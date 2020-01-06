import React, { useEffect } from 'react';
import { Switch, Route, useHistory } from 'react-router';
import { useDispatch, useSelector } from 'react-redux';
import { BrowserRouter } from 'react-router-dom';
import { CssBaseline, createMuiTheme, Button } from '@material-ui/core';
import { ThemeProvider } from '@material-ui/core/styles';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer, toast, cssTransition } from 'react-toastify';

import { IState } from '../types/redux';
import HarmonySocket from '../socket/socket';
import { AppDispatch } from '../redux/store';
import { SetPrimary, SetSecondary, InvertTheme, SetInputStyle, SetSelf } from '../redux/AppReducer';

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
	const dispatch = useDispatch<AppDispatch>();
	const { theme } = useSelector((state: IState) => state);
	useSocketHandler(harmonySocket, useHistory());
	useEffect(() => {
		const localPrimary = localStorage.getItem('primary');
		const localSecondary = localStorage.getItem('secondary');
		const localType = localStorage.getItem('themetype');
		const localInputStyle = localStorage.getItem('inputstyle');
		const localSelf = localStorage.getItem('self');
		if (localPrimary) {
			dispatch(SetPrimary(JSON.parse(localPrimary)));
		}
		if (localSecondary) {
			dispatch(SetSecondary(JSON.parse(localSecondary)));
		}
		if (localType !== theme.type) {
			dispatch(InvertTheme());
		}
		if (localInputStyle === 'standard' || localInputStyle === 'filled' || localInputStyle === 'outlined') {
			dispatch(SetInputStyle(localInputStyle));
		}
		if (localSelf) {
			dispatch(SetSelf(JSON.parse(localSelf)));
		}
	}, []);

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
