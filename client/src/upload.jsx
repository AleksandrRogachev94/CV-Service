import React, { useEffect } from 'react';
import axios from 'axios';

const Upload = () => {
  useEffect(() => {
    (async () => {
      const res = await axios.get('/api/health');
      console.log(res);
    })()
  }, [])
  return (
    <h1>Upload Page</h1>
  );
};

export default Upload;