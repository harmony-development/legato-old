import React, { useRef, useState } from 'react';
import { useStyles } from './styles';
import { Tooltip, IconButton, TextField, Box } from '@material-ui/core';
import { Warning, Image } from '@material-ui/icons';
import FileCard from './FileCard/FileCard';
import { Events } from '../../../../socket/socket';
import { socketServer } from '../../../Root';

interface IProps {
    name: string;
}

const ChatBox: React.FC<IProps> = (props: IProps) => {
    const classes = useStyles();
    const [draftMessage, setDraftMessage] = useState('');
    const inputFile = useRef<HTMLInputElement>(null);
    const [fileQueue, setFileQueue] = useState<string[]>([]);

    const sendFile = (): void => {
        if (inputFile.current) {
            inputFile.current.click();
        }
    };

    const onFileSelected = (event: React.ChangeEvent<HTMLInputElement>): void => {
        if (event.currentTarget.files && event.currentTarget.files.length > 0) {
            Array.from(event.currentTarget.files).forEach((file) => {
                if (file.type.startsWith('image/') && file.size < 33554432) {
                    const imageReader = new FileReader();
                    imageReader.readAsDataURL(file);
                    imageReader.addEventListener('load', () => {
                        if (typeof imageReader.result === 'string') {
                            setFileQueue((prevQueue) => [...prevQueue, imageReader.result as string]);
                        }
                    });
                }
            });
        }
    };

    const chatBoxKeyEvent = (event: React.KeyboardEvent<HTMLInputElement>): void => {
        if (event.key === 'Enter') {
            event.preventDefault();
            if (!event.shiftKey) {
                const message = (event.target as HTMLInputElement).value;
                socketServer.connection.emit(Events.MESSAGE, {
                    token: localStorage.getItem('token'),
                    message,
                    files: fileQueue
                });
                setFileQueue([]);
                (event.target as HTMLInputElement).selectionEnd = 0;
                setDraftMessage('');
            } else {
                setDraftMessage(draftMessage + '\n');
            }
        }
    };

    const removeFromQueue = (index: number): void => {
        setFileQueue([...fileQueue.slice(0, index), ...fileQueue.slice(index + 1)]);
    };

    return (
        <>
            <Box display='flex' className={classes.fileQueue}>
                {fileQueue.map((file: string, index) => {
                    return <FileCard image={file} removeFromQueue={removeFromQueue} index={index} key={index} />;
                })}
            </Box>
            <div className={classes.messageBoxContainer}>
                <div className={classes.valign}>
                    <input type='file' id='file' multiple ref={inputFile} style={{ display: 'none' }} onChange={onFileSelected} />
                    <Tooltip title='Send Image'>
                        <IconButton onClick={sendFile}>
                            <Image />
                        </IconButton>
                    </Tooltip>
                </div>
                <TextField onKeyPress={chatBoxKeyEvent} value={draftMessage} onChange={(event): void => setDraftMessage(event.target.value)} className={classes.chatBox} label={socketServer.connection.connected ? 'Send Message' : 'Currently Offline'} multiline rows='2' margin='normal' />
                {!socketServer.connection.connected ? (
                    <div className={classes.valign}>
                        <Tooltip title='Currently Offline. You will not be able to send messages.'>
                            <IconButton>
                                <Warning />
                            </IconButton>
                        </Tooltip>
                    </div>
                ) : (
                    undefined
                )}
            </div>
        </>
    );
};

export default ChatBox;
