import React from 'react';
import { AppBar, Toolbar, Typography } from '@material-ui/core';

export const App = () => {
    return (
        <div>
            <AppBar position='static'>
                <Toolbar>
                    <Typography variant='h6'>Harmony</Typography>
                </Toolbar>
            </AppBar>
        </div>
    );
};
