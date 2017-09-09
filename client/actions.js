export const SET_CONFIG = 'SET_CONFIG';
export const LOAD_TALKS = 'LOAD_TALKS';
export const SET_TALKS = 'SET_TALKS';
export const ERROR_TALKS = 'ERROR_TALKS';
export const LOAD_POPULAR_TALKS = 'LOAD_POPULAR_TALKS';
export const SET_POPULAR_TALKS = 'SET_POPULAR_TALKS';
export const ERROR_POPULAR_TALKS = 'ERROR_POPULAR_TALKS';

export function setConfig(config) {
  return { type: SET_CONFIG, config };
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
