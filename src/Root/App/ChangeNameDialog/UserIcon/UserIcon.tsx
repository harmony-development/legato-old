import React, { useRef, Dispatch, SetStateAction } from 'react';
import { Avatar, Button } from '@material-ui/core';
import { useStyles } from './styles';

interface IProps {
  icon: string;
  setIcon: Dispatch<SetStateAction<string>>;
}

const UserIcon: React.FC<IProps> = (props: IProps) => {
  const inputFile = useRef<HTMLInputElement>(null);
  const classes = useStyles();

  const onFileSelected = (event: React.ChangeEvent<HTMLInputElement>): void => {
    if (event.currentTarget.files && event.currentTarget.files.length > 0) {
      const newIcon = event.currentTarget.files[0];
      if (newIcon.type.startsWith('image/') && newIcon.size < 33554432) {
        const imageReader = new FileReader();
        imageReader.readAsDataURL(newIcon);
        imageReader.addEventListener('load', () => {
          if (typeof imageReader.result === 'string') {
            props.setIcon(imageReader.result);
          }
        });
      }
    }
  };

  const onIconButtonPressed = (): void => {
    if (inputFile.current) {
      inputFile.current.click();
    }
  };

  return (
    <div className={classes.iconRoot}>
      <Avatar src={props.icon} />
      <input type='file' ref={inputFile} style={{ display: 'none' }} onChange={onFileSelected} />
      <Button variant='contained' color='primary' className={classes.changeIconButton} onClick={onIconButtonPressed}>
        Change Icon
      </Button>
    </div>
  );
};

export default UserIcon;
