import React, { Component } from 'react';
import Helmet from 'react-helmet';
import Nav from '../nav';
import Footer from '../footer';

export default class App extends Component {

  render() {
    return <div>
      <Helmet title='Codetalks' />
      <Nav />
      <div className="app-content container is-fluid">
        {this.props.children}
      </div>
      <Footer />
    </div>;
  }

}
