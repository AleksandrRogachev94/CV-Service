import React, { useState } from 'react';
import axios from 'axios';

import RecognitionResult from './recognition-result';

const UploadFile = ({ }) => {
  const [imagesDict, setImagesDict] = useState(null);

  const uploadFile = async ({ target: { files } }) => {
    const file = files[0];
    try {
      const { data: { bucket, key } } = await axios.post('/api/upload', file, {
        headers: {
          'Content-Type': 'image/jpeg'
        }
      });
      const { data } = await axios.post('/api/recognitions', { bucket, key });
      // const data = { "Bicycle": [{ "conf": 81.72513580322266, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Bicycle-81.73/0-81.73.png" }, { "conf": 64.36577606201172, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Bicycle-81.73/1-64.37.png" }], "Car": [{ "conf": 88.21635437011719, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Car-93.81/2-88.22.png" }, { "conf": 78.5203628540039, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Car-93.81/3-78.52.png" }, { "conf": 62.126853942871094, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Car-93.81/4-62.13.png" }, { "conf": 93.80887603759766, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Car-93.81/0-93.81.png" }, { "conf": 91.24565124511719, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Car-93.81/1-91.25.png" }], "Person": [{ "conf": 77.71991729736328, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Person-99.73/9-77.72.png" }, { "conf": 97.10074615478516, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Person-99.73/5-97.10.png" }, { "conf": 97.70856475830078, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Person-99.73/4-97.71.png" }, { "conf": 97.81217956542969, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Person-99.73/3-97.81.png" }, { "conf": 95.72144317626953, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Person-99.73/7-95.72.png" }, { "conf": 99.69076538085938, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Person-99.73/1-99.69.png" }, { "conf": 95.92339324951172, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Person-99.73/6-95.92.png" }, { "conf": 93.2274398803711, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Person-99.73/8-93.23.png" }, { "conf": 99.72705078125, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Person-99.73/0-99.73.png" }, { "conf": 99.67610168457031, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Person-99.73/2-99.68.png" }], "Wheel": [{ "conf": 65.10760498046875, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Wheel-85.14/1-65.11.png" }, { "conf": 85.14473724365234, "url": "https://go-cvservice-assets.s3.amazonaws.com/d3a26ce1-7a67-4434-88c7-da70858dfd67.jpg-results/Wheel-85.14/0-85.14.png" }] }
      setImagesDict(data);
      console.log('recognize: ', data)
    } catch (err) {
      console.log(err);
      alert('Something went wrong. Try again later');
    }
  }

  return (
    <div>
      <div>
        <label htmlFor="file">Choose file to upload</label>
        <input id="file" type="file" accept="image/jpeg" onChange={uploadFile} />
      </div>

      {imagesDict && <RecognitionResult key={imagesDict} results={imagesDict} />}
    </div>
  );
};

export default UploadFile;