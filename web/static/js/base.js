// import { notesList } from "./data.js";

export default class NotesList {
  constructor() {
    this.list = [];
  }

  main() {
    this.listenRefreshButton();
    this.getList();
  }

  listenRefreshButton() {
    const $fresh_button = $(".fresh-button>button");
    $fresh_button.on("click", () => {
      this.getList();
    });
  }

  getNote(id) {
    const url = `http://localhost:8080/note/${id}`;
    // console.log(url);

    return new Promise((resolve) => {
      $.ajax({
        url,
        type: "get",
        success(resp) {
          if (resp.message === "success") {
            resolve(resp.note);
          }
        },
        error(err) {
          console.log(err);
        },
      });
    });
  }

  async getList() {
    // console.log("get list");
    const tasks = [];
    for (let i = 1; i <= 10; i++) {
      tasks.push(this.getNote(i));
      // console.log(note);
    }

    // 并行发送请求
    this.list = await Promise.all(tasks);
    // console.log(this.list);
    if (!this.list instanceof Array) {
      console.error("request note error.");
    }

    // 按时间降序排序
    this.list = this.list.sort(
      (a, b) => new Date(b.modified_time) - new Date(a.modified_time)
    );

    this.renderList();
  }

  renderList() {
    const $notesListBody = $(".notes-list-body");
    $notesListBody.empty(); // 先清除所有子元素

    // console.log(this.list);

    this.list.forEach((item, idx) => {
      const $temp = $(`
    <div class="list-item">
          <div class="item-left">${idx + 1}.</div>
          <div class="item-right">
            <div class="item-right-body">
              <div class="item-body-title">${item.title}</div>
              <div class="item-body-content">${item.content}</div>
            </div>
            <div class="item-footer">last_modified: ${item.modified_time}</div>
          </div>
    </div>`);
      $notesListBody.append($temp);
    });
  }
}
