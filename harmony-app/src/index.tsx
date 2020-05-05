import './index.css';
import './Root/ContextMenu.css';

import React from 'react';
import ReactDOM from 'react-dom';
import { CircularProgress } from '@material-ui/core';
import { ToastContainer, cssTransition } from 'react-toastify';
import { BrowserRouter } from 'react-router-dom';
import { PersistGate } from 'redux-persist/integration/react';
import { Provider } from 'react-redux';

import { store, persistor } from './redux/store';
import * as serviceWorker from './serviceWorker';
import { ConfirmationContextProvider } from './Root/App/ConfirmationContext';
import { Root } from './Root/Root';
import './i18n/i18n';

const Index = React.memo(() => {
	return (
		<Provider store={store}>
			<PersistGate loading={<CircularProgress />} persistor={persistor}>
				<ToastContainer
					position="bottom-left"
					pauseOnFocusLoss={false}
					transition={cssTransition({
						enter: 'zoomIn',
						exit: 'slideOut',
						duration: 200,
					})}
				/>
				<ConfirmationContextProvider>
					<BrowserRouter>
						<Root />
					</BrowserRouter>
				</ConfirmationContextProvider>
			</PersistGate>
		</Provider>
	);
});

ReactDOM.render(<Index />, document.getElementById('root'));

serviceWorker.register();
