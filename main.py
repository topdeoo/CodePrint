import win32print
import tempfile
import win32api
import time
from flask import Flask, render_template, request
import csv

app = Flask(__name__)


@app.route('/', methods=['GET'])
def home():
    return render_template('home.html')


@app.route('/upload', methods=['POST'])
def upload():
    team_name, password, context = request.form['team_name'], request.form['team_pwd'], request.form['team_code']
    checker = db.get(team_name)
    if checker[0] == password and len(context) <= 20000:
        # filename = team_name + time.strftime("%H:%M:%S", time.localtime(time.time())) + ".txt"
        # content = context + "\n\n" + team_name + "\t" + checker[1]
        # open(filename, "w+").write(content)
        filename = tempfile.mktemp(".txt")
        open(filename, "w").write("TEAM_NAME :  " + team_name + "\n" +
                                  context + "\n" +
                                  "TEAM_NAME :  " + team_name + "\n" +
                                  "LOCATION :  " + checker[1])
        win32api.ShellExecute(
            0,
            "printto",
            filename,
            '"%s"' % win32print.GetDefaultPrinter(),
            ".",
            0
        )
        return render_template('success.html')
    else:
        return render_template('failed.html')


# request 'team_info.csv' format: team_name, password
if __name__ == "__main__":
    db = {}
    with open("team_info.csv", "r") as f:
        cr = csv.DictReader(f)
        for row in cr:
            db[row['team_name']] = (row['password'], row['location'])
    app.run(host="0.0.0.0", port="8080")
