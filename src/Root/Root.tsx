/**
 * This file is intended for adding things such as Redux Providers and other things.
 */
import React from 'react';
import { Provider } from 'react-redux';
import { store } from '../store/store';
import App from './App/App';

const Root: React.FC<{}> = () => {
  return (
    <Provider store={store}>
      <App />
    </Provider>
  );
};

export default Root;
