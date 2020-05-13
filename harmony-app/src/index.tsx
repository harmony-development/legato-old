import './index.css';
import './ContextMenu.css';

import React from 'react';
import ReactDOM from 'react-dom';
import { CircularProgress } from '@material-ui/core';
import { ToastContainer, cssTransition } from 'react-toastify';
import { BrowserRouter } from 'react-router-dom';
import { Provider } from 'react-redux';

import { store } from './redux/store';
import * as serviceWorker from './serviceWorker';
import { Root } from './Root';
import './i18n/i18n';

const Index = React.memo(() => {
	return (
		<Provider store={store}>
			<ToastContainer
				position="bottom-left"
				pauseOnFocusLoss={false}
				transition={cssTransition({
					enter: 'zoomIn',
					exit: 'slideOut',
					duration: 200,
				})}
			/>
			<BrowserRouter>
				<Root />
			</BrowserRouter>
		</Provider>
	);
});

ReactDOM.render(<Index />, document.getElementById('root'));

serviceWorker.register();
