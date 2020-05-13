import React, { useEffect } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import axios from 'axios';

import Navbar from './navbar';
import Footer from './footer';
import About from './about';
import Upload from './upload';
import Login from './login';
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

  useEffect(() => {
    (async () => {
      try {
        const res = await axios.get('/api/health');
        console.log(res);
      } catch (err) {
        console.error(err);
        alert('Server is currently down. Please try again later');
      }
    })();
  }, []);

  return (
    <Router>
      <Box className={classes.pageContainer}>
        <Navbar />
        <Container maxWidth="lg" className={classes.contentContainer}>
          <Paper className={classes.contentPaper}>
            <Switch>
              <Route exact path="/" component={Upload} />
              <Route exact path="/about" component={About} />
              <Route exact path="/login" component={Login} />
            </Switch>
          </Paper>
        </Container>
        <Footer />
      </Box>
    </Router>
  );
}

export default App;
