// partially derived from https://dev.to/dmtrkovalenko/the-neatest-way-to-handle-alert-dialogs-in-react-1aoe

import React from 'react';

import { CommonDialog } from './CommonDialog';

interface DialogOptions {
	title: string;
	description: string;
	type: 'confirm' | 'alert';
}

interface Props {
	children: JSX.Element;
}

const CommonDialogContext = React.createContext<(options: DialogOptions) => Promise<void>>(Promise.reject);

export const CommonDialogContextProvider = ({ children }: Props) => {
	const [dialogState, setConfirmState] = React.useState<DialogOptions | null>(null);
	const pendingDialogRef = React.useRef<{
		resolve: () => void;
		reject: () => void;
	}>();

	const openDialog = (options: DialogOptions) => {
		setConfirmState(options);
		return new Promise<void>((resolve, reject) => {
			pendingDialogRef.current = { resolve, reject };
		});
	};

	const cancelHandler = () => {
		if (pendingDialogRef.current) {
			pendingDialogRef.current.reject();
		}
		setConfirmState(null);
	};

	const confirmHandler = () => {
		if (pendingDialogRef.current) {
			pendingDialogRef.current.resolve();
		}
		setConfirmState(null);
	};

	return (
		<>
			<CommonDialogContext.Provider value={openDialog}>{children}</CommonDialogContext.Provider>
			{dialogState ? (
				<CommonDialog onSubmit={confirmHandler} onClose={cancelHandler} open={Boolean(dialogState)} {...dialogState} />
			) : (
				undefined
			)}
		</>
	);
};

export const useDialog = () => React.useContext(CommonDialogContext);
