import {
  loadTalk,
  setTalk,
  errorTalk,
  loadTalkProfile,
  setTalkProfile,
  errorTalkProfile,
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
  const qs = q ? `?q=${encodeURIComponent(q)}` : '';
  return fetch('/api/v1/talk' + qs)
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
    .catch(err => store.dispatch(errorPopularTalks(err)))
    .catch(() => store.dispatch(loadPopularTalks(false)));
}

export function getTalk(store, params) {
  const { talkId } = params;
  store.dispatch(loadTalk(true));
  return fetch(`/api/v1/talk/${talkId}`)
    .then(res => res.json())
    .then((talk) => {
      store.dispatch(setTalk(talk));
    })
    .catch(err => store.dispatch(errorTalk(err)))
    .finally(() => store.dispatch(loadTalk(false)));
}

export function getTalkProfile(store, params) {
  const { talkId } = params;
  store.dispatch(loadTalkProfile(true));
  return fetch(`/api/v1/talk/${talkId}/profile`)
    .then(res => res.json())
    .then((profile) => {
      store.dispatch(setTalkProfile(profile));
    })
    .catch(err => store.dispatch(errorTalkProfile(err)))
    .finally(() => store.dispatch(loadTalkProfile(false)));
}
