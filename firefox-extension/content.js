// Hide the comments section
const hideComments = () => {
    const commentsSection = document.getElementById('comments');
    if (commentsSection) {
      commentsSection.style.display = 'none';
    }
  };

  document.addEventListener('DOMContentLoaded', hideComments);

  const observer = new MutationObserver(hideComments);
  observer.observe(document.body, { childList: true, subtree: true });
  