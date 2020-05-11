import React, { useState } from 'react';
import PropTypes from 'prop-types';
import Box from '@material-ui/core/Box';
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';
import InputLabel from '@material-ui/core/InputLabel';
import FormControl from '@material-ui/core/FormControl';
import Paper from '@material-ui/core/Paper';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  selectLabel: {
    margin: '1em auto',
    padding: '1em',
    width: '30%',
    [theme.breakpoints.down('sm')]: {
      width: '80%'
    }
  },
  imageResult: {
    marginTop: '1em',
    height: '200px',
    '& img': {
      height: '100%'
    }
  }
}));

const RecognitionResult = ({ results }) => {
  const classes = useStyles();
  const labels = Object.keys(results);
  const [selectedLabel, setSelectedLabel] = useState(labels[0]);
  console.log(results, selectedLabel)
  const handleChange = ({ target: { value } }) => {
    setSelectedLabel(value);
  };

  return (
    <Box>
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
        <>
          <h4>{selectedLabel}</h4>
          {results[selectedLabel].map(image => (
            <div key={image.url} className={classes.imageResult} >
              <a target="_blank" href={image.url}>
                <img src={image.url} />
              </a>
            </div>
          ))}
        </>
      )}
    </Box>
  );
}

RecognitionResult.propTypes = {
  results: PropTypes.objectOf(PropTypes.arrayOf(PropTypes.shape({
    conf: PropTypes.number.isRequired,
    url: PropTypes.string.isRequired,
  }))).isRequired,
};

export default RecognitionResult;