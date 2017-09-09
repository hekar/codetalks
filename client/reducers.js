import { combineReducers } from 'redux';
import {
  SET_CONFIG,
  LOAD_TALKS,
  SET_TALKS,
  ERROR_TALKS,
  LOAD_POPULAR_TALKS,
  SET_POPULAR_TALKS,
  ERROR_POPULAR_TALKS
} from './actions';

function config(state = {}, action) {
  switch (action.type) {
  case SET_CONFIG:
    return action.config;
  default:
    return state;
  }
}

function talks(state = {}, action) {
  switch (action.type) {
  case LOAD_TALKS:
    return action.talksLoading;
  case SET_TALKS:
    return action.talks;
  case ERROR_TALKS:
    return action.error;
  default:
    return state;
  }
}

function popularTalks(state ={}, action) {
  switch (action.type) {
  case LOAD_POPULAR_TALKS:
    return action.popularTalksLoading;
  case SET_POPULAR_TALKS:
    return action.popularTalks;
  case ERROR_POPULAR_TALKS:
    return action.popularTalksError;
  default:
    return state;
  }
}

export default combineReducers({ config, talks, popularTalks });
