import React, { Component } from 'react';
import Helmet from 'react-helmet';
import { Link } from 'react-router-dom';
import Grid from '../grid';

export default class Talks extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  /*eslint-disable */
  static onEnter({store, nextState, replaceState, callback}) {
    // Load here any data.
    callback(); // this call is important, don't forget it
  }
  /*eslint-enable */

  componentWillMount() {
    fetch('/api/v1/talk')
      .then(res => res.json())
      .then(({ talks }) => this.setState({ talks }))
      .catch(err => this.setState({ err }));
  }

  render() {
    debugger;
    const { err, talks } = this.state;

    if (err) {
      return <div>Error: {err}</div>;
    } else if (talks) {
      return <div>
        <Grid></Grid>
      </div>;
    } else {
      return <div>Loading..</div>;
    }
  }

}
