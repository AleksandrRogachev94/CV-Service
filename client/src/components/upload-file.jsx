import React from 'react';
import PropTypes from 'prop-types';

import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import AddIcon from '@material-ui/icons/Add';
import CircularProgress from '@material-ui/core/CircularProgress';

const useStyles = makeStyles((theme) => ({
  input: {
    display: 'none'
  },
  uploadButton: {
    textAlign: 'center',
    padding: theme.spacing(3)
  }
}));

const UploadFile = ({ handleUpload, isLoading }) => {
  const classes = useStyles();

  return (
    <div>
      <div className={classes.uploadButton}>
        <input
          accept="image/jpeg"
          className={classes.input}
          id="raised-button-file"
          type="file"
          onChange={handleUpload}
          disabled={isLoading}
        />
        <label htmlFor="raised-button-file">
          <Button
            size="large"
            variant="outlined"
            component="span"
            startIcon={<AddIcon />}
            disabled={isLoading}
          >
            {isLoading ? <CircularProgress /> : 'Upload New File'}
          </Button>
        </label>
      </div>
    </div>
  );
};

UploadFile.propTypes = {
  handleUpload: PropTypes.func.isRequired,
  isLoading: PropTypes.bool,
};

UploadFile.defaultProps = {
  isLoading: false,
};

export default UploadFile;