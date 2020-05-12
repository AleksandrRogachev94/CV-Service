import React, { useState } from 'react';
import axios from 'axios';

import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import AddIcon from '@material-ui/icons/Add';
import RecognitionResult from './recognition-result';

const useStyles = makeStyles((theme) => ({
  input: {
    display: 'none'
  },
  uploadButton: {
    textAlign: 'center',
    padding: theme.spacing(3)
  }
}));

const UploadFile = ({ }) => {
  const classes = useStyles();
  const [imagesDict, setImagesDict] = useState(null);
  const [sourceUrl, setSourceUrl] = useState(null);

  const uploadFile = async ({ target: { files } }) => {
    const file = files[0];
    try {
      const { data: { bucket, key } } = await axios.post('/api/upload', file, {
        headers: {
          'Content-Type': 'image/jpeg'
        }
      });
      setSourceUrl("https://" + bucket + ".s3.amazonaws.com/" + key);
      const { data } = await axios.post('/api/recognitions', { bucket, key });
      // const data = { "Person": [{ "conf": 53.552207946777344, "url": "https://go-cvservice-assets.s3.amazonaws.com/c02f9e38-c66d-41d3-a6d3-dc47292fcc40.jpg-results/Person-99.91/6-53.55.png" }, { "conf": 56.04245376586914, "url": "https://go-cvservice-assets.s3.amazonaws.com/c02f9e38-c66d-41d3-a6d3-dc47292fcc40.jpg-results/Person-99.91/5-56.04.png" }, { "conf": 67.2090072631836, "url": "https://go-cvservice-assets.s3.amazonaws.com/c02f9e38-c66d-41d3-a6d3-dc47292fcc40.jpg-results/Person-99.91/4-67.21.png" }, { "conf": 99.58563232421875, "url": "https://go-cvservice-assets.s3.amazonaws.com/c02f9e38-c66d-41d3-a6d3-dc47292fcc40.jpg-results/Person-99.91/2-99.59.png" }, { "conf": 99.21111297607422, "url": "https://go-cvservice-assets.s3.amazonaws.com/c02f9e38-c66d-41d3-a6d3-dc47292fcc40.jpg-results/Person-99.91/3-99.21.png" }, { "conf": 99.85400390625, "url": "https://go-cvservice-assets.s3.amazonaws.com/c02f9e38-c66d-41d3-a6d3-dc47292fcc40.jpg-results/Person-99.91/0-99.85.png" }, { "conf": 99.81889343261719, "url": "https://go-cvservice-assets.s3.amazonaws.com/c02f9e38-c66d-41d3-a6d3-dc47292fcc40.jpg-results/Person-99.91/1-99.82.png" }] }
      setImagesDict(data);
    } catch (err) {
      console.log(err);
      alert('Something went wrong. Try again later');
    }
  }

  return (
    <div>
      <div className={classes.uploadButton}>
        <input
          accept="image/jpeg"
          className={classes.input}
          id="raised-button-file"
          type="file"
          onChange={uploadFile}
        />
        <label htmlFor="raised-button-file">
          <Button
            size="large"
            variant="outlined"
            component="span"
            startIcon={<AddIcon />}
          >
            Upload New File
          </Button>
        </label>
      </div>

      {imagesDict && <RecognitionResult key={imagesDict} results={imagesDict} sourceUrl={sourceUrl} />}
    </div>
  );
};

export default UploadFile;