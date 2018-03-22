import React, { Component } from 'react';
import Helmet from 'react-helmet';
import Nav from '../nav';
import Footer from '../footer';

export default class App extends Component {

  render() {
    return <div>
      <Helmet title='Codetalks'>
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" />
      </Helmet>
      <Nav />
      <div className="app-content">
        {this.props.children}
      </div>
      <Footer />
    </div>;
  }

}
