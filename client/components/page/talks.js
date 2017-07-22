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
    const { err, talks } = this.state;

    if (err) {
      return <div>Error: {err}</div>;
    } else if (talks) {
      const rows = talks.map(talk =>{
        const tags = talk.Tags.map(tag => <div><a href={'/tag/' + tag}>tag</a></div>);
        return (
          <div key={talk.TalkID}>
            <a href={talk.Url}><img src={talk.ThumbnailUrl} /></a>
            <a href={talk.Url}>{talk.Name}</a>
            {tags}
            {JSON.stringify(talk)}
          </div>
        );
      });
      return <div>
        {rows}
      </div>;
    } else {
      return <div>Loading..</div>;
    }
  }

}
