import React, { Component } from 'react';
import { Link } from 'react-router-dom';

export default class Footer extends Component {
  render() {
    return (
      <footer className="footer">
        <div className="columns">
          <div className="column is-2">
            <strong>Codetalks</strong>
            <ul>
              <li><a href="">Terms</a></li>
              <li><a href="">Privacy</a></li>
              <li><a href="">About</a></li>
            </ul>
          </div>
          <div className="column is-2">
            <strong>Codetalks</strong>
            <ul>
              <li><a href="">Terms</a></li>
              <li><a href="">Privacy</a></li>
              <li><a href="">About</a></li>
            </ul>
          </div>
          <div className="column is-2">
            <strong>Codetalks</strong>
            <ul>
              <li><a href="">Terms</a></li>
              <li><a href="">Privacy</a></li>
              <li><a href="">About</a></li>
            </ul>
          </div>
        </div>
      </footer>
    );
  }
}