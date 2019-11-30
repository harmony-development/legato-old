import React from 'react';
import { Dialog, DialogTitle, DialogContent, CardActionArea } from '@material-ui/core';

interface IProps {
    open: boolean;
    setOpen: (newval: boolean) => void;
    image: string;
}

const ImageDialog: React.FC<IProps> = (props: IProps) => {
    return (
        <Dialog open={props.open} onClose={(): void => props.setOpen(false)}>
            <DialogTitle>Image Preview</DialogTitle>
            <DialogContent>
                <CardActionArea>
                    <img src={props.image} style={{ width: '100%' }} alt='' />
                </CardActionArea>
            </DialogContent>
        </Dialog>
    );
};

export default ImageDialog;
