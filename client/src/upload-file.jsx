import React from 'react';
import axios from 'axios';

const UploadFile = ({ }) => {

  const uploadFile = async ({ target: { files } }) => {
    const file = files[0];
    try {
      const { data: { bucket, key } } = await axios.post('/api/upload', file, {
        headers: {
          'Content-Type': 'image/jpeg'
        }
      });
      const { data } = await axios.post('/api/recognitions', { bucket, key });
      console.log('recognize: ', data)
    } catch (err) {
      console.log(err);
      alert('Something went wrong. Try again later');
    }


  }

  return (
    <div>
      <label htmlFor="file">Choose file to upload</label>
      <input id="file" type="file" accept="image/jpeg" onChange={uploadFile} />
    </div>

  )
};

export default UploadFile;