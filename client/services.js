import {
  loadTalks,
  setTalks,
  errorTalks
} from './actions';

export function searchTalks(store, params) {
  store.dispatch(loadTalks(true));
  return fetch('/api/v1/talk')
    .then(res => res.json())
    .then(({ talks }) => {
      console.log(JSON.stringify(talks));
      store.dispatch(setTalks(talks));
    })
    .catch(err => store.dispatch(errorTalks(err)));
}
