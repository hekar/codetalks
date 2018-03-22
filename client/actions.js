export const SET_CONFIG = 'SET_CONFIG';
export const LOAD_TALK = 'LOAD_TALK';
export const SET_TALK = 'SET_TALK';
export const ERROR_TALK = 'ERROR_TALK';
export const LOAD_TALK_PROFILE = 'LOAD_TALK_PROFILE';
export const SET_TALK_PROFILE = 'SET_TALK_PROFILE';
export const ERROR_TALK_PROFILE = 'ERROR_TALK_PROFILE';
export const LOAD_TALKS = 'LOAD_TALKS';
export const SET_TALKS = 'SET_TALKS';
export const ERROR_TALKS = 'ERROR_TALKS';
export const LOAD_POPULAR_TALKS = 'LOAD_POPULAR_TALKS';
export const SET_POPULAR_TALKS = 'SET_POPULAR_TALKS';
export const ERROR_POPULAR_TALKS = 'ERROR_POPULAR_TALKS';
export const LOAD_RECENTLY_ADDED_TALKS = 'LOAD_RECENTLY_ADDED_TALKS';
export const SET_RECENTLY_ADDED_TALKS = 'SET_RECENTLY_ADDED_TALKS';
export const ERROR_RECENTLY_ADDED_TALKS = 'ERROR_RECENTLY_ADDED_TALKS';

export function setConfig(config) {
  return { type: SET_CONFIG, config };
}

export function loadTalk(talkLoading) {
  return { type: LOAD_TALK, talkLoading };
}

export function setTalk(talk) {
  return { type: SET_TALK, talk };
}

export function errorTalk(error) {
  return { type: ERROR_TALK, error };
}

export function loadTalkProfile(loading) {
  return { type: LOAD_TALK_PROFILE, loading };
}

export function setTalkProfile(profile) {
  return { type: SET_TALK_PROFILE, profile };
}

export function errorTalkProfile(error) {
  return { type: ERROR_TALK_PROFILE, error };
}

export function loadTalks(talksLoading) {
  return { type: LOAD_TALKS, talksLoading };
}

export function setTalks(talks) {
  return { type: SET_TALKS, talks };
}

export function errorTalks(error) {
  return { type: ERROR_TALKS, error };
}

export function loadPopularTalks(popularTalksLoading) {
  return { type: LOAD_POPULAR_TALKS, popularTalksLoading };
}

export function setPopularTalks(popularTalks) {
  return { type: SET_POPULAR_TALKS, popularTalks };
}

export function errorPopularTalks(popularTalksError) {
  return { type: ERROR_POPULAR_TALKS, popularTalksError };
}

export function loadRecentlyAddedTalks(recentlyAddedTalksLoading) {
  return { type: LOAD_RECENTLY_ADDED_TALKS, recentlyAddedTalksLoading };
}

export function setRecentlyAddedTalks(recentlyAddedTalks) {
  return { type: SET_RECENTLY_ADDED_TALKS, recentlyAddedTalks };
}

export function errorRecentlyAddedTalks(recentlyAddedTalksError) {
  return { type: ERROR_RECENTLY_ADDED_TALKS, recentlyAddedTalksError };
}