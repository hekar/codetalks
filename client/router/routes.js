import React from 'react';
import { Route, IndexRoute, Redirect } from 'react-router';
import App from '#app/components/app';
import Homepage from '#app/components/page/home';
import Usage from '#app/components/page/usage';
import NotFound from '#app/components/page/not-found';
import Search from '#app/components/page/search';
import Register from '#app/components/page/register';
import Talks from '#app/components/page/talks';
import Signin from '../components/page/signin';

const disableServerSideRendering = true;

/**
 * Returns configured routes for different
 * environments. `w` - wrapper that helps skip
 * data fetching with onEnter hook at first time.
 * @param {Object} - any data for static loaders and first-time-loading marker
 * @returns {Object} - configured routes
 */
export default ({store, first}) => {

  // Make a closure to skip first request
  function w(loader) {
    return (nextState, replaceState, callback) => {
      if (!disableServerSideRendering && first.time) {
        first.time = false;
        return callback();
      }
      return loader ? loader({
        store,
        nextState,
        replaceState,
        callback
      }) : callback();
    };
  }

  return (
    <App>
      <Route path="/" component={App}>
        <IndexRoute component={Homepage}
          onEnter={w(Homepage.onEnter)}/>
        <Route path="/usage" component={Usage}
          onEnter={w(Usage.onEnter)}/>
        <Route path="/search" component={Search}
          onEnter={w(Search.onEnter)}/>
        <Route path="/register" component={Register}
          onEnter={w(Register.onEnter)}/>
        <Route path="/signin" component={Signin}
          onEnter={w(Signin.onEnter)}/>
        <Route path="/talks" component={Talks}
          onEnter={w(Talks.onEnter)}/>
        <Route path="*" component={NotFound}
          onEnter={w(NotFound.onEnter)}/>
      </Route>
    </App>
  );
};
