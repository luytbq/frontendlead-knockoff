function renderQuestionTags(question, parrentEl) {
    try {
        const difficultyArr = getDifficultyArray(question);
        const difficulty = difficultyArr[0];
        const label = document.createElement('label');
        label.classList.add('question-tag');
        label.classList.add('difficulty');
        label.classList.add(difficulty.toLowerCase());
        label.innerText = difficulty;
        parrentEl.appendChild(label);

        const renderType = (type) => {
            const label2 = document.createElement('label');
            label2.classList.add('question-tag');
            label.classList.add(type.toLowerCase());
            label2.innerText = type;
            parrentEl.appendChild(label2);
        }
        const types = getTypeArray(question);
        types.forEach(renderType);


        if (types.some(e => e.toLowerCase() !== 'coding')) {
            const span = document.createElement('label');
            span.classList.add('error');
            span.innerText = 'This question type is currently not supported'
            parrentEl.appendChild(span)
        }
    } catch (err) {
        console.error(err);
    }
}

function getDifficultyArray(question) {
    return question.common?.difficulty || question.difficulty || ['Difficulty unknown'];
}

function getTypeArray(question) {
    const types = question.info?.type || question.type || ['Type unknown'];
    if (Array.isArray(types)) return types;
    return [types];
}