import React from 'react';

import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';

const About = () => {
  return (
    <Box>
      <Typography variant="h3">Computer Vision Service App</Typography>

      <Box mt={2}>
        <Typography>
          This app helps users to extract useful information from images using computer vision algorithms.
          As of right now, the only feature is entities extraction. The app identifies entities in the image and extracts them as separate donwloadable images.
          See more details on the project github page: <a href="https://github.com/AleksandrRogachev94/CV-Service">https://github.com/AleksandrRogachev94/CV-Service</a>
          <br /><br />
          This app is being implemented by Aleksandr Rogachev
          <br />
          Email: aleksandr.rogachev1994@gmail.com
        </Typography>
      </Box>

    </Box>
  );
};

export default About;