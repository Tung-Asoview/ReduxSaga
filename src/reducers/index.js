import { combineReducers } from 'redux';
import foods from './foods';
import artists from './artists';

const allReducers = combineReducers({
    foods , artists
});

export default allReducers;