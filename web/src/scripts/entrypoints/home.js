import '../../styles/entrypoints/home.scss';
import './base.js';
import Parallax from 'parallax-js'

document.addEventListener("DOMContentLoaded", () => {
  let scene = document.getElementById('scene');
  let parallax = new Parallax(scene);
});
