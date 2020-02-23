// partially derived from https://dev.to/dmtrkovalenko/the-neatest-way-to-handle-alert-dialogs-in-react-1aoe

import React from 'react';

import { ConfirmDialog } from './Dialog/ConfirmDialog/ConfirmDialog';

interface ConfirmOptions {
	title: string;
	description: string;
}

interface Props {
	children: JSX.Element;
}

const ConfirmationContext = React.createContext<(options: ConfirmOptions) => Promise<void>>(Promise.reject);

export const ConfirmationContextProvider = ({ children }: Props) => {
	const [confirmState, setConfirmState] = React.useState<ConfirmOptions | null>(null);
	const pendingDialogRef = React.useRef<{
		resolve: () => void;
		reject: () => void;
	}>();

	const openConfirmation = (options: ConfirmOptions) => {
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
			<ConfirmationContext.Provider value={openConfirmation}>{children}</ConfirmationContext.Provider>
			{confirmState ? (
				<ConfirmDialog
					onSubmit={confirmHandler}
					onClose={cancelHandler}
					open={Boolean(confirmState)}
					{...confirmState}
				/>
			) : (
				undefined
			)}
		</>
	);
};

export const useConfirmationContext = () => React.useContext(ConfirmationContext);
