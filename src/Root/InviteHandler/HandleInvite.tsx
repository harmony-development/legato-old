import React, { useEffect, useState } from 'react';
import { useParams, useHistory } from 'react-router';
import { harmonySocket } from '../Root';
import { Typography, Button } from '@material-ui/core';
import { useInviteHandlerStyle } from './InviteHandlerStyle';

export const InviteHandler = () => {
    const { invitecode } = useParams();
    const history = useHistory();
    const [errorMessage, setErrorMessage] = useState<string | null>(null);
    const classes = useInviteHandlerStyle();

    useEffect(() => {
        harmonySocket.events.addListener('joinguild', (raw: any) => {
            console.log(raw);
            if (!raw['message']) {
                setErrorMessage(null);
                history.push('/app');
            } else {
                setErrorMessage(raw['message']);
            }
        });
        harmonySocket.events.addListener('open', () => {
            if (invitecode) {
                harmonySocket.joinGuild(invitecode);
            }
        });
    }, [history, invitecode]);

    return (
        <div>
            {errorMessage ? (
                <div className={classes.errorRoot}>
                    <Typography variant='h1' align='center' className={classes.errorMsg}>
                        404
                        <br />
                        {errorMessage}
                    </Typography>
                    <Button variant='contained' color='secondary' className={classes.errorBtn} onClick={() => history.push('/')}>
                        Return To Login
                    </Button>
                </div>
            ) : (
                undefined
            )}
        </div>
    );
};
