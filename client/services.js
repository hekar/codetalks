import {
  loadTalks,
  setTalks,
  errorTalks,
  loadPopularTalks,
  setPopularTalks,
  errorPopularTalks
} from './actions';

export function searchTalks(store, params) {
  store.dispatch(loadTalks(true));
  const queryParams = new URLSearchParams(document.location.search.substring(1));
  const q = queryParams.get('q');

  return fetch('/api/v1/talk?q=' + encodeURIComponent(q))
    .then(res => res.json())
    .then(({ talks }) => {
      store.dispatch(setTalks(talks));
    })
    .catch(err => store.dispatch(errorTalks(err)));
}

export function popularTalks(store, params) {
  store.dispatch(loadPopularTalks(true));
  return fetch('/api/v1/talk/popular')
    .then(res => res.json())
    .then(({ talks }) => {
      console.log('talks', talks);
      store.dispatch(setPopularTalks(talks));
    })
    .catch(err => store.dispatch(errorPopularTalks(err)));
}
