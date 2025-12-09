filename = "out/daily_reflect.tex"
start_pg = 9
is_leap = False

# l-page (even page number): trim=1.3cm 2cm 0.3cm 1.35cm
# r-page (odd page number): trim=0.22cm 2cm 1.3cm 1.35cm

# Important strings
l_page = "1.3cm 2cm 0.3cm 1.35cm"
r_page = "0.22cm 2cm 1.3cm 1.35cm"
add_page_string = r"""

\vspace*{-15.37cm}
\makebox[14.93cm][r]{%
  \raisebox{-\totalheight}[0pt][0pt]{
    \includegraphics[page=<page>,trim=<trim>,clip,width=5cm]{wotaw.pdf}
  }
}"""
magic_string = (
    r"\myUnderline{Daily log}"
    + "\n"
    + r"\myMash{\myNumDailyDiaryLog}{\myNumDotWidthFull}"
    + "\n"
)
magic_len = len(magic_string)


with open(filename, "r", encoding="utf-8") as file:
    tex = file.read()

pos = -1
pg = start_pg
while True:
    pos = tex.find(magic_string, pos + 1)
    if pos == -1:
        break
    pos += magic_len - 1
    include_str = add_page_string.replace("<page>", str(pg)).replace(
        "<trim>", l_page if pg % 2 == 0 else r_page
    )
    tex = tex[:pos] + include_str + tex[pos:]
    if pg == 67 and not is_leap:
        pg += 1
    pg += 1

with open(filename, "w", encoding="utf-8") as file:
    file.write(tex)
