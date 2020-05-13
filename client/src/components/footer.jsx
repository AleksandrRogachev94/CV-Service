import React from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';

const useStyles = makeStyles((theme) => ({
  footer: {
    height: '2.5rem',
    width: '100%',
    // paddingTop: '1rem',
    position: 'absolute',
    bottom: 0,
    backgroundColor: theme.palette.grey[300],
  },
  content: {
    marginTop: '0.5rem',
    marginLeft: '1rem'
  }
}));

const Footer = () => {
  const classes = useStyles();
  const year = new Date().getFullYear();

  return (
    <Box className={classes.footer}>
      <Typography className={classes.content} variant="body2">&copy; Copyright {year}, Aleksandr Rogachev</Typography>
    </Box>
  );
}

export default Footer;