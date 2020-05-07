import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import Navbar from './navbar';
import About from './about';
import Upload from './upload';
import { makeStyles } from '@material-ui/core/styles';

import Container from '@material-ui/core/Container';
import Paper from '@material-ui/core/Paper';

const useStyles = makeStyles((theme) => ({
  appContainer: {
    height: '100%',
    marginTop: theme.spacing(2),
    padding: theme.spacing(2),
  }
}));

function App() {
  const classes = useStyles();

  return (
    <Router>
      <Navbar />
      <Container maxWidth="lg">
        <Paper className={classes.appContainer}>
          <Switch>
            <Route exact path="/" component={Upload} />
            <Route exact path="/about" component={About} />
          </Switch>
        </Paper>
      </Container>
    </Router>
  );
}

export default App;
