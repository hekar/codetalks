import React, { Component } from 'react';
import { connect } from 'react-redux';
import Helmet from 'react-helmet';
import { Link } from 'react-router-dom';
import { setConfig } from '../../actions';

class Usage extends Component {

  /*eslint-disable */
  static onEnter({store, nextState, replaceState, callback}) {
    fetch('/api/v1/conf').then((r) => {
      return r.json();
    }).then((conf) => {
      store.dispatch(setConfig(conf));
      console.log('Faked connection latency! Please, take a look ---> `server/api.go:22`');
      callback();
    });
  }
  /*eslint-enable */

  render() {
    return <div>
      <Helmet title='Usage' />
      <h2>Usage:</h2>
      <div>
        <span>// TODO: write an article</span>
        <pre>config:
          {JSON.stringify(this.props.config, null, 2)}</pre>
      </div>
      <br />
      go <Link to='/'>home</Link>
    </div>;
  }

}

export default connect(store => ({ config: store.config }))(Usage);
