export const SET_CONFIG = 'SET_CONFIG';
export const LOAD_TALKS = 'LOAD_TALKS';
export const SET_TALKS = 'SET_TALKS';

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
