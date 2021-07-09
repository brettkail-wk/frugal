package com.workiva.frugal.server;

import java.util.AbstractMap;
import java.util.AbstractSet;
import java.util.Collections;
import java.util.Enumeration;
import java.util.Iterator;
import java.util.List;
import java.util.Map;
import java.util.Set;

abstract class AbstractServletRequestHeaders extends AbstractMap<String, List<String>> {
    protected abstract Enumeration<String> names();
    protected abstract Enumeration<String> values(String name);

    @Override
    public List<String> get(Object key) {
        if (!(key instanceof String)) {
            return null;
        }

        Enumeration<String> en = values((String) key);
        return en.hasMoreElements() ? Collections.list(en) : null;
    }

    @Override
    public Set<Map.Entry<String, List<String>>> entrySet() {
        return new AbstractSet<Map.Entry<String, List<String>>>() {
            @Override
            public Iterator<Map.Entry<String, List<String>>> iterator() {
                Enumeration<String> en = names();
                return new Iterator<Map.Entry<String, List<String>>>() {
                    @Override
                    public boolean hasNext() {
                        return en.hasMoreElements();
                    }

                    @Override
                    public Map.Entry<String, List<String>> next() {
                        String name = en.nextElement();
                        return new SimpleEntry<>(name, get(name));
                    }
                };
            }

            @Override
            public int size() {
                int size = 0;
                for (Enumeration<String> en = names(); en.hasMoreElements(); en.nextElement()) {
                    size++;
                }
                return size;
            }
        };
    }
}
