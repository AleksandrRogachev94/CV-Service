import React, { useState } from 'react';
import axios from 'axios';

import Typography from '@material-ui/core/Typography';
import UploadFile from './upload-file'
import RecognitionResult from './recognition-result';

const Upload = () => {
  const [imagesDict, setImagesDict] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const [sourceUrl, setSourceUrl] = useState(null);

  const uploadFile = async ({ target: { files } }) => {
    const file = files[0];
    try {
      setIsLoading(true);
      const { data: { bucket, key } } = await axios.post('/api/upload', file, {
        headers: {
          'Content-Type': 'image/jpeg'
        }
      });
      const { data } = await axios.post('/api/recognitions', { bucket, key });
      // const data = { "Person": [{ "conf": 53.552207946777344, "url": "https://go-cvservice-assets.s3.amazonaws.com/c02f9e38-c66d-41d3-a6d3-dc47292fcc40.jpg-results/Person-99.91/6-53.55.png" }, { "conf": 56.04245376586914, "url": "https://go-cvservice-assets.s3.amazonaws.com/c02f9e38-c66d-41d3-a6d3-dc47292fcc40.jpg-results/Person-99.91/5-56.04.png" }, { "conf": 67.2090072631836, "url": "https://go-cvservice-assets.s3.amazonaws.com/c02f9e38-c66d-41d3-a6d3-dc47292fcc40.jpg-results/Person-99.91/4-67.21.png" }, { "conf": 99.58563232421875, "url": "https://go-cvservice-assets.s3.amazonaws.com/c02f9e38-c66d-41d3-a6d3-dc47292fcc40.jpg-results/Person-99.91/2-99.59.png" }, { "conf": 99.21111297607422, "url": "https://go-cvservice-assets.s3.amazonaws.com/c02f9e38-c66d-41d3-a6d3-dc47292fcc40.jpg-results/Person-99.91/3-99.21.png" }, { "conf": 99.85400390625, "url": "https://go-cvservice-assets.s3.amazonaws.com/c02f9e38-c66d-41d3-a6d3-dc47292fcc40.jpg-results/Person-99.91/0-99.85.png" }, { "conf": 99.81889343261719, "url": "https://go-cvservice-assets.s3.amazonaws.com/c02f9e38-c66d-41d3-a6d3-dc47292fcc40.jpg-results/Person-99.91/1-99.82.png" }] }
      setImagesDict(data);
      setSourceUrl("https://" + bucket + ".s3.amazonaws.com/" + key);
    } catch (err) {
      console.log(err);
      alert('Something went wrong. Try again later');
    } finally {
      setIsLoading(false);
    }
  }

  return (
    <>
      <Typography variant="h3" style={{ textAlign: 'center' }}>Process New Image</Typography>
      <UploadFile handleUpload={uploadFile} isLoading={isLoading} />
      {imagesDict && sourceUrl &&
        <RecognitionResult key={sourceUrl} results={imagesDict} sourceUrl={sourceUrl} />}
    </>
  );
};

export default Upload;