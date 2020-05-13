import React, { useState } from 'react';
import PropTypes from 'prop-types';
import Box from '@material-ui/core/Box';
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';
import InputLabel from '@material-ui/core/InputLabel';
import FormControl from '@material-ui/core/FormControl';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  selectLabel: {
    textAlign: 'center',
    margin: '1em auto',
    padding: '1em',
    width: '30%',
    [theme.breakpoints.down('sm')]: {
      width: '80%'
    }
  },
  selectedLabel: {
    textAlign: 'center'
  },
  sourceImage: {
    margin: theme.spacing(2),
    textAlign: 'center',
    '& img': {
      width: '50%',
    }
  },

  resultItem: {
    marginTop: '1em',
    textAlign: 'center',
    margin: '0 auto',
  },
  imageResult: {
    height: '250px',
    '& img': {
      height: '250px'
    }
  },
}));

const RecognitionResult = ({ results, sourceUrl }) => {
  const classes = useStyles();
  const labels = results && Object.keys(results);
  const [selectedLabel, setSelectedLabel] = useState(labels[0]);
  const handleChange = ({ target: { value } }) => {
    setSelectedLabel(value);
  };

  if (!results || Object.keys(results).length <= 0) {
    return (
      <Box>
        <Typography>No Results</Typography>
      </Box>
    )
  }

  return (
    <Box>
      {sourceUrl && (
        <Box className={classes.sourceImage}>
          <Typography variant="h5">Source Image</Typography>
          <img src={sourceUrl} alt="source" />
        </Box>
      )}

      <Paper className={classes.selectLabel}>
        <FormControl>
          <InputLabel id="select-label">Label</InputLabel>
          <Select
            labelId="select-label"
            id="select-label"
            value={selectedLabel}
            onChange={handleChange}
          >
            {labels.map(label => (
              <MenuItem key={label} value={label}>{label}</MenuItem>
            ))}
          </Select>
        </FormControl>
      </Paper>

      {results && selectedLabel && (
        results[selectedLabel].map(image => (
          <Box key={image.url} className={classes.resultItem} >
            <Typography variant="subtitle1">Confidence: {image.conf.toFixed(2)}%</Typography>
            <a href={image.url} className={classes.imageResult}>
              <img src={image.url} alt="result" />
            </a>
          </Box>
        ))
      )}
    </Box>
  );
}

RecognitionResult.propTypes = {
  results: PropTypes.objectOf(PropTypes.arrayOf(PropTypes.shape({
    conf: PropTypes.number.isRequired,
    url: PropTypes.string.isRequired,
  }))).isRequired,
  sourceUrl: PropTypes.string.isRequired,
};

export default RecognitionResult;