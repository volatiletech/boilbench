#!/usr/bin/env python
# Python 2 script, sorry :(

import glob
import os
import errno
import plotly.plotly as plotly
import plotly.graph_objs as graphing

def make_graph(filename, title, ytitle, string_data):
    lines = string_data.splitlines()
    x = []
    y = []
    err = []
    for i in lines:
        # Don't take min/max values
        name, _, _, average, error = i.split(',')
        x.append(name)
        y.append(float(average))
        err.append(float(error))

    colors = [ 'rgb(49,171,95)', 'rgb(49, 110, 171)', 'rgb(212, 109, 57)',
            'rgb(148, 62, 154)', 'rgb(54, 176, 165)']
    if len(lines) > 5:
        colors.append('rgb(184, 75, 75)')

    trace = graphing.Bar(
        x = x,
        y = y,
        error_y = dict(
            type  = 'data',
            array = err,
            visible = True
        ),
        marker = dict(
            color = colors
        )
    )

    layout = graphing.Layout(
        title = title,
        autosize = False,
        width = 300,
        height = 300,
        margin = graphing.Margin(
           l = 60, r = 30, b = 60, t = 40, pad = 4 
        ),
        xaxis = graphing.XAxis(
            title = "ORM"
        ),
        yaxis = graphing.YAxis(
            title = ytitle
        ),
        bargap = 5
    )

    data = graphing.Data([trace])
    fig = graphing.Figure(data=data, layout=layout)
    plotly.image.save_as(fig, filename = filename)

title_table = dict(nsop = 'Speed', bop = 'Memory', aop = 'Allocations')
yaxis_table = dict(nsop = 'ns/op', bop = 'B/op', aop = 'allocs/op')

# Python2 sucks :(
try:
    os.makedirs('graphs')
except OSError as exc:
    if exc.errno != errno.EEXIST or not os.path.isdir('graphs'):
        raise

for kind in ['Delete', 'Insert', 'RawBind', 'SelectAll', 'SelectComplex', 'SelectSubset', 'Update']:
    for bench in ['nsop', 'bop', 'aop']:
        with open('graph_data/%s_%s.csv' % (kind, bench)) as f:
            data = f.read().strip()
            make_graph(
                    'graphs/%s_%s.png' % (kind.lower(), bench),
                    '%s %s' % (kind, title_table[bench]),
                    yaxis_table[bench],
                    data
            )
