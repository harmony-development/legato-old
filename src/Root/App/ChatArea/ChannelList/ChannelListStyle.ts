import { makeStyles } from '@material-ui/styles';
import { Theme } from '@material-ui/core';

export const useChannelListStyle = makeStyles((theme: Theme) => ({
    selectedChannel: {
        backgroundColor: theme.palette.type === 'dark' ? theme.palette.grey[800] : theme.palette.grey[400]
    },
    nested: {
        paddingLeft: theme.spacing(4)
    },
    newChannelInput: {
        marginLeft: theme.spacing(2),
        marginRight: theme.spacing(2)
    }
}));
