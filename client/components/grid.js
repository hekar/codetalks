import React, { Component } from 'react';

export default class Grid extends Component {

  static get defaultProps() {
    return {
      headers: [{
        text: ''
      }],
      rows: [{
        cells: [{
          key: '',
          widget: 'anchor',
          href: '',
          text: ''
        }]
      }]
    };
  }

  render() {
    const headers = this.props.headers.map(header => (
      <th>{header.text}</th>
    ));

    const rows = this.props.rows.map(row => (
      <tr>{
        row.cells.map(cell => {
          switch (cell.widget) {
          case 'anchor':
            return (
              <td key={cell.key}>
                <a href={cell.href}>{cell.text}</a>
              </td>
            );
          default:
            return (
              <td key={cell.key}>{cell.text}</td>
            );
          }
        })
      }
      </tr>
    ));
    
    return (
      <table>
        <caption></caption>
        <thead>
          <tr>
            {headers}
          </tr>
        </thead>
        <tbody>{rows}</tbody>
      </table>
    );
  }
}
