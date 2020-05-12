import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import Navbar from './navbar';
import Footer from './footer';
import About from './about';
import Upload from './upload';
import { makeStyles } from '@material-ui/core/styles';

import Container from '@material-ui/core/Container';
import Paper from '@material-ui/core/Paper';
import Box from '@material-ui/core/Box';

const useStyles = makeStyles((theme) => ({
  pageContainer: {
    position: 'relative', // for the footer
    minHeight: '100vh'
  },
  contentContainer: {
    height: '100%',
    padding: theme.spacing(3),
    paddingBottom: '4rem',
  },
  contentPaper: {
    padding: theme.spacing(2),
  }
}));

function App() {
  const classes = useStyles();

  return (
    <Router>
      <Box className={classes.pageContainer}>
        <Navbar />
        <Container maxWidth="lg" className={classes.contentContainer}>
          <Paper className={classes.contentPaper}>
            <Switch>
              <Route exact path="/" component={Upload} />
              <Route exact path="/about" component={About} />
            </Switch>
          </Paper>
        </Container>
        <Footer />
      </Box>
    </Router>
  );
}

export default App;
