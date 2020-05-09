import React, { useEffect } from 'react';
import axios from 'axios';

import Typography from '@material-ui/core/Typography';
import UploadFile from './upload-file'

const Upload = () => {
  useEffect(() => {
    (async () => {
      const res = await axios.get('/api/health');
      console.log(res);
    })()
  }, [])
  return (
    <>
      <Typography variant="h3">Upload New Image</Typography>
      <UploadFile />
    </>
  );
};

export default Upload;