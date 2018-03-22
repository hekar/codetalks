import React, { Component } from 'react';

export default class Pagination extends Component {

  static get defaultProps() {
    return {
      currentPage: 1,
      pageCount: 99
    };
  }

  render() {
    const {
      currentPage,
      pageCount
    } = this.props;
    const pages = new Array(pageCount).fill(0).map((_, i) => {
      const index = i + 1;
      const className = "pagination-link" + ((currentPage === index) ? " is-current" : "");
      return (<li key={'' + index}><a className={className} aria-label={'Page ' + index} aria-current="page">{index}</a></li>);
    });

    const moreThanTen = pages.length > 10;
    if (moreThanTen) {
      const marker = 4;
      const separator = (<span className="pagination-ellipsis">&hellip;</span>);
      if (currentPage + marker < pages.length) {
        pages.splice(currentPage + marker, pages.length - 2 - marker);
      }
      if (currentPage - marker > 1) {
        pages.splice(2, currentPage - marker);
      }
      if (currentPage > 10) {
        pages.splice(1, 0, (<li key="-1">{separator}</li>));
      }
      pages.splice(pages.length - 1, 0, (<li key="-2">{separator}</li>));
    }
    return (
      <nav className="pagination" role="navigation" aria-label="pagination">
        <ul className="pagination-list">
          {pages}
        </ul>
        <a className="pagination-previous">Previous</a>
        <a className="pagination-next">Next page</a>
      </nav>
    );
  }
}
