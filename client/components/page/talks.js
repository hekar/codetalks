import React, { Component } from 'react';
import Helmet from 'react-helmet';
import { Link } from 'react-router-dom';
import Grid from '../grid';

export default class Talks extends Component {
  /*eslint-disable */
  static onEnter({store, nextState, replaceState, callback}) {
    // Load here any data.
    callback(); // this call is important, don't forget it
  }
  /*eslint-enable */

  render() {
    
    return (
      <div>
        <h1>fadsdf</h1>
        <Grid></Grid>
      </div>
    );
  }

}
